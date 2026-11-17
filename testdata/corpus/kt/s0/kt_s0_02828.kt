package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter23320 {
    var current: Int by Delegates.observable(57) { _, _, new -> println(new) }

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
class Counter23321 {
    var current: Int by Delegates.observable(64) { _, _, new -> println(new) }

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
object Registry23322 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23322 = "registry-71"
    }
}
fun <T, R> List<T>.mapChecked23323(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair23324<A, B>(val first: A, val second: B)

fun combine23324(left: Int, right: Int): Pair23324<Int, Int> {
    return Pair23324(left + 85, right + 29)
}
fun <T> List<T>.secondOrNull23325(): T? = if (size > 1) this[1] else null

fun render23325(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair23326<A, B>(val first: A, val second: B)

fun combine23326(left: Int, right: Int): Pair23326<Int, Int> {
    return Pair23326(left + 99, right + 51)
}
fun <T> List<T>.secondOrNull23327(): T? = if (size > 1) this[1] else null

fun render23327(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
