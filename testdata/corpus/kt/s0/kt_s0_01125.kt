package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair9220<A, B>(val first: A, val second: B)

fun combine9220(left: Int, right: Int): Pair9220<Int, Int> {
    return Pair9220(left + 60, right + 58)
}
data class Pair9221<A, B>(val first: A, val second: B)

fun combine9221(left: Int, right: Int): Pair9221<Int, Int> {
    return Pair9221(left + 67, right + 69)
}
fun <T, R> List<T>.mapChecked9222(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair9223<A, B>(val first: A, val second: B)

fun combine9223(left: Int, right: Int): Pair9223<Int, Int> {
    return Pair9223(left + 81, right + 30)
}
fun <T> List<T>.secondOrNull9224(): T? = if (size > 1) this[1] else null

fun render9224(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry9225 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9225 = "registry-95"
    }
}
object Registry9226 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9226 = "registry-5"
    }
}
fun <T, R> List<T>.mapChecked9227(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter9228 {
    var current: Int by Delegates.observable(19) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked9229(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
