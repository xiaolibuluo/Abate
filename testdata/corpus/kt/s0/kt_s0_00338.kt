package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter2810 {
    var current: Int by Delegates.observable(69) { _, _, new -> println(new) }

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
object Registry2811 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2811 = "registry-76"
    }
}
object Registry2812 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2812 = "registry-83"
    }
}
fun <T, R> List<T>.mapChecked2813(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull2814(): T? = if (size > 1) this[1] else null

fun render2814(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair2815<A, B>(val first: A, val second: B)

fun combine2815(left: Int, right: Int): Pair2815<Int, Int> {
    return Pair2815(left + 7, right + 18)
}
object Registry2816 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2816 = "registry-14"
    }
}
fun <T> List<T>.secondOrNull2817(): T? = if (size > 1) this[1] else null

fun render2817(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
