package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull14970(): T? = if (size > 1) this[1] else null

fun render14970(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair14971<A, B>(val first: A, val second: B)

fun combine14971(left: Int, right: Int): Pair14971<Int, Int> {
    return Pair14971(left + 73, right + 68)
}
class Counter14972 {
    var current: Int by Delegates.observable(80) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull14973(): T? = if (size > 1) this[1] else null

fun render14973(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter14974 {
    var current: Int by Delegates.observable(94) { _, _, new -> println(new) }

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
data class Pair14975<A, B>(val first: A, val second: B)

fun combine14975(left: Int, right: Int): Pair14975<Int, Int> {
    return Pair14975(left + 101, right + 51)
}
object Registry14976 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14976 = "registry-11"
    }
}
