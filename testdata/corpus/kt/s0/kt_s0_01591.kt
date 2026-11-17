package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked13160(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair13161<A, B>(val first: A, val second: B)

fun combine13161(left: Int, right: Int): Pair13161<Int, Int> {
    return Pair13161(left + 48, right + 21)
}
