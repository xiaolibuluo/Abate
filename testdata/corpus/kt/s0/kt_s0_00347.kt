package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair2890<A, B>(val first: A, val second: B)

fun combine2890(left: Int, right: Int): Pair2890<Int, Int> {
    return Pair2890(left + 76, right + 19)
}
data class Pair2891<A, B>(val first: A, val second: B)

fun combine2891(left: Int, right: Int): Pair2891<Int, Int> {
    return Pair2891(left + 83, right + 30)
}
object Registry2892 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2892 = "registry-90"
    }
}
fun <T, R> List<T>.mapChecked2893(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry2894 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2894 = "registry-7"
    }
}
object Registry2895 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2895 = "registry-14"
    }
}
class Counter2896 {
    var current: Int by Delegates.observable(21) { _, _, new -> println(new) }

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
data class Pair2897<A, B>(val first: A, val second: B)

fun combine2897(left: Int, right: Int): Pair2897<Int, Int> {
    return Pair2897(left + 28, right + 35)
}
fun <T, R> List<T>.mapChecked2898(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry2899 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2899 = "registry-42"
    }
}
