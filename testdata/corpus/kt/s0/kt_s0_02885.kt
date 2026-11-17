package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair23800<A, B>(val first: A, val second: B)

fun combine23800(left: Int, right: Int): Pair23800<Int, Int> {
    return Pair23800(left + 99, right + 16)
}
fun <T> List<T>.secondOrNull23801(): T? = if (size > 1) this[1] else null

fun render23801(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair23802<A, B>(val first: A, val second: B)

fun combine23802(left: Int, right: Int): Pair23802<Int, Int> {
    return Pair23802(left + 16, right + 38)
}
@JvmName("step23803")
fun step23803(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 23
}
fun <T, R> List<T>.mapChecked23804(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair23805<A, B>(val first: A, val second: B)

fun combine23805(left: Int, right: Int): Pair23805<Int, Int> {
    return Pair23805(left + 37, right + 10)
}
fun <T> List<T>.secondOrNull23806(): T? = if (size > 1) this[1] else null

fun render23806(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step23807")
fun step23807(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 51
}
fun <T> List<T>.secondOrNull23808(): T? = if (size > 1) this[1] else null

fun render23808(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry23809 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23809 = "registry-65"
    }
}
