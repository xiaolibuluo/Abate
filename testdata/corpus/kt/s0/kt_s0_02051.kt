package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked16950(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step16951")
fun step16951(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 28
}
fun <T> List<T>.secondOrNull16952(): T? = if (size > 1) this[1] else null

fun render16952(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter16953 {
    var current: Int by Delegates.observable(42) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
data class Pair16954<A, B>(val first: A, val second: B)

fun combine16954(left: Int, right: Int): Pair16954<Int, Int> {
    return Pair16954(left + 49, right + 23)
}
object Registry16955 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16955 = "registry-56"
    }
}
@JvmName("step16956")
fun step16956(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 63
}
fun <T, R> List<T>.mapChecked16957(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair16958<A, B>(val first: A, val second: B)

fun combine16958(left: Int, right: Int): Pair16958<Int, Int> {
    return Pair16958(left + 77, right + 67)
}
