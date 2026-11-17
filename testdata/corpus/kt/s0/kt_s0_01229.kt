package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair10160<A, B>(val first: A, val second: B)

fun combine10160(left: Int, right: Int): Pair10160<Int, Int> {
    return Pair10160(left + 21, right + 45)
}
fun <T> List<T>.secondOrNull10161(): T? = if (size > 1) this[1] else null

fun render10161(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull10162(): T? = if (size > 1) this[1] else null

fun render10162(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked10163(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull10164(): T? = if (size > 1) this[1] else null

fun render10164(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull10165(): T? = if (size > 1) this[1] else null

fun render10165(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair10166<A, B>(val first: A, val second: B)

fun combine10166(left: Int, right: Int): Pair10166<Int, Int> {
    return Pair10166(left + 63, right + 50)
}
fun <T, R> List<T>.mapChecked10167(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step10168")
fun step10168(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 77
}
