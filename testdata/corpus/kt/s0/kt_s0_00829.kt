package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair6820<A, B>(val first: A, val second: B)

fun combine6820(left: Int, right: Int): Pair6820<Int, Int> {
    return Pair6820(left + 44, right + 25)
}
