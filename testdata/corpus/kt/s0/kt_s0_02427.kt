package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair19960<A, B>(val first: A, val second: B)

fun combine19960(left: Int, right: Int): Pair19960<Int, Int> {
    return Pair19960(left + 54, right + 12)
}
