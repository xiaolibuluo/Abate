package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair11690<A, B>(val first: A, val second: B)

fun combine11690(left: Int, right: Int): Pair11690<Int, Int> {
    return Pair11690(left + 70, right + 18)
}
