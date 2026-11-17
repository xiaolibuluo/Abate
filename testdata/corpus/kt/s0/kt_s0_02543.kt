package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair20920<A, B>(val first: A, val second: B)

fun combine20920(left: Int, right: Int): Pair20920<Int, Int> {
    return Pair20920(left + 41, right + 13)
}
