package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair8770<A, B>(val first: A, val second: B)

fun combine8770(left: Int, right: Int): Pair8770<Int, Int> {
    return Pair8770(left + 57, right + 48)
}
