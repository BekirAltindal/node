/** Node-type implemeted in Go language
 *
 * @file
 * @author Steffen Vogel <svogel2@eonerc.rwth-aachen.de>
 * @copyright 2014-2022, Institute for Automation of Complex Power Systems, EONERC
 * @license Apache 2.0
 *********************************************************************************/

#pragma once

typedef void *_go_plugin;
typedef void *_go_plugin_list;
typedef void *_go_logger;

typedef  void (*_go_register_node_factory_cb)(_go_plugin_list pl, char *name, char *desc, int flags);
typedef void (*_go_logger_log_cb)(_go_logger l, int level, char *msg);

_go_logger * _go_logger_get(char *name);

void _go_logger_log(_go_logger *l, int level, char *msg);
