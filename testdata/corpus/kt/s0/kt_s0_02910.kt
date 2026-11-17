package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair24020<A, B>(val first: A, val second: B)

fun combine24020(left: Int, right: Int): Pair24020<Int, Int> {
    return Pair24020(left + 94, right + 48)
}
data class Pair24021<A, B>(val first: A, val second: B)

fun combine24021(left: Int, right: Int): Pair24021<Int, Int> {
    return Pair24021(left + 101, right + 59)
}
fun <T> List<T>.secondOrNull24022(): T? = if (size > 1) this[1] else null

fun render24022(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step24023")
fun step24023(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 18
}
fun <T, R> List<T>.mapChecked24024(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
