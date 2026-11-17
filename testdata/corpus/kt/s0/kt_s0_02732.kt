package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step22550")
fun step22550(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 26
}
@JvmName("step22551")
fun step22551(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
data class Pair22552<A, B>(val first: A, val second: B)

fun combine22552(left: Int, right: Int): Pair22552<Int, Int> {
    return Pair22552(left + 40, right + 17)
}
fun <T, R> List<T>.mapChecked22553(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair22554<A, B>(val first: A, val second: B)

fun combine22554(left: Int, right: Int): Pair22554<Int, Int> {
    return Pair22554(left + 54, right + 39)
}
