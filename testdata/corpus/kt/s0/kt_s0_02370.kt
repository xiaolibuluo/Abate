package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair19470<A, B>(val first: A, val second: B)

fun combine19470(left: Int, right: Int): Pair19470<Int, Int> {
    return Pair19470(left + 96, right + 35)
}
