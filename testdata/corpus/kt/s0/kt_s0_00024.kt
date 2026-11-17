package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step200")
fun step200(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 71
}
data class Pair201<A, B>(val first: A, val second: B)

fun combine201(left: Int, right: Int): Pair201<Int, Int> {
    return Pair201(left + 78, right + 38)
}
data class Pair202<A, B>(val first: A, val second: B)

fun combine202(left: Int, right: Int): Pair202<Int, Int> {
    return Pair202(left + 85, right + 49)
}
fun <T> List<T>.secondOrNull203(): T? = if (size > 1) this[1] else null

fun render203(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair204<A, B>(val first: A, val second: B)

fun combine204(left: Int, right: Int): Pair204<Int, Int> {
    return Pair204(left + 99, right + 10)
}
@JvmName("step205")
fun step205(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 9
}
