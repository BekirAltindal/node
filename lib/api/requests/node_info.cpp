/** The "node" API ressource.
 *
 * @author Steffen Vogel <svogel2@eonerc.rwth-aachen.de>
 * @copyright 2014-2022, Institute for Automation of Complex Power Systems, EONERC
 * @license GNU General Public License (version 3)
 *
 * VILLASnode
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *********************************************************************************/

#include <jansson.h>
#include <uuid/uuid.h>

#include <villas/super_node.hpp>
#include <villas/node.hpp>
#include <villas/utils.hpp>
#include <villas/stats.hpp>
#include <villas/api/session.hpp>
#include <villas/api/requests/node.hpp>
#include <villas/api/response.hpp>

namespace villas {
namespace node {
namespace api {

class NodeInfoRequest : public NodeRequest {

public:
	using NodeRequest::NodeRequest;

	virtual Response * execute()
	{
		if (method != Session::Method::GET)
			throw InvalidMethod(this);

		if (body != nullptr)
			throw BadRequest("Nodes endpoint does not accept any body data");

		auto *json_node = node->toJson();

		auto sigs = node->getOutputSignals();
		if (sigs) {
			auto *json_out = json_object_get(json_node, "out");
			if (json_out)
				json_object_set_new(json_out, "signals", sigs->toJson());
		}

		return new JsonResponse(session, HTTP_STATUS_OK, json_node);
	}
};

/* Register API request */
static char n[] = "node";
static char r[] = "/node/(" RE_NODE_NAME "|" RE_UUID ")";
static char d[] = "retrieve info of a node";
static RequestPlugin<NodeInfoRequest, n, r, d> p;

} /* namespace api */
} /* namespace node */
} /* namespace villas */
