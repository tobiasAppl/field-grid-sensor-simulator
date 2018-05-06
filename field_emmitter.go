/**
 * Author: Tobias Appl <tobias.appl@gmail.com>
**/

package main;

type field_emmitter_2d interface {
    calculate_field_effect(target_pos Point2d) float64
}
