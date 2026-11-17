package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair19000<A, B>(val first: A, val second: B)

fun combine19000(left: Int, right: Int): Pair19000<Int, Int> {
    return Pair19000(left + 67, right + 11)
}
data class Pair19001<A, B>(val first: A, val second: B)

fun combine19001(left: Int, right: Int): Pair19001<Int, Int> {
    return Pair19001(left + 74, right + 22)
}
