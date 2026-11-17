package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry15830 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15830 = "registry-20"
    }
}
fun <T, R> List<T>.mapChecked15831(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair15832<A, B>(val first: A, val second: B)

fun combine15832(left: Int, right: Int): Pair15832<Int, Int> {
    return Pair15832(left + 34, right + 10)
}
data class Pair15833<A, B>(val first: A, val second: B)

fun combine15833(left: Int, right: Int): Pair15833<Int, Int> {
    return Pair15833(left + 41, right + 21)
}
class Counter15834 {
    var current: Int by Delegates.observable(48) { _, _, new -> println(new) }

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
class Counter15835 {
    var current: Int by Delegates.observable(55) { _, _, new -> println(new) }

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
