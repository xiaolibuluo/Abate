package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair2790<A, B>(val first: A, val second: B)

fun combine2790(left: Int, right: Int): Pair2790<Int, Int> {
    return Pair2790(left + 43, right + 10)
}
object Registry2791 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2791 = "registry-50"
    }
}
fun <T> List<T>.secondOrNull2792(): T? = if (size > 1) this[1] else null

fun render2792(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked2793(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair2794<A, B>(val first: A, val second: B)

fun combine2794(left: Int, right: Int): Pair2794<Int, Int> {
    return Pair2794(left + 71, right + 54)
}
class Counter2795 {
    var current: Int by Delegates.observable(78) { _, _, new -> println(new) }

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
class Counter2796 {
    var current: Int by Delegates.observable(85) { _, _, new -> println(new) }

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
