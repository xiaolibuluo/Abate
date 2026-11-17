package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair1250<A, B>(val first: A, val second: B)

fun combine1250(left: Int, right: Int): Pair1250<Int, Int> {
    return Pair1250(left + 78, right + 30)
}
fun <T> List<T>.secondOrNull1251(): T? = if (size > 1) this[1] else null

fun render1251(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry1252 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1252 = "registry-92"
    }
}
fun <T, R> List<T>.mapChecked1253(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull1254(): T? = if (size > 1) this[1] else null

fun render1254(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter1255 {
    var current: Int by Delegates.observable(16) { _, _, new -> println(new) }

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
