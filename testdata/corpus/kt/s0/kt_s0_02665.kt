package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22020<A, B>(val first: A, val second: B)

fun combine22020(left: Int, right: Int): Pair22020<Int, Int> {
    return Pair22020(left + 16, right + 51)
}
data class Pair22021<A, B>(val first: A, val second: B)

fun combine22021(left: Int, right: Int): Pair22021<Int, Int> {
    return Pair22021(left + 23, right + 62)
}
fun <T, R> List<T>.mapChecked22022(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step22023")
fun step22023(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 37
}
