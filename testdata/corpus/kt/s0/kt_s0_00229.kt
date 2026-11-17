package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked1890(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome1891 {
    data class Success(val payload: String) : Outcome1891()
    data object Failure : Outcome1891()
}

fun describe1891(outcome: Outcome1891): String = when (outcome) {
    is Outcome1891.Success -> outcome.payload
    Outcome1891.Failure -> "failed"
}
data class Pair1892<A, B>(val first: A, val second: B)

fun combine1892(left: Int, right: Int): Pair1892<Int, Int> {
    return Pair1892(left + 51, right + 12)
}
@JvmName("step1893")
fun step1893(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 58
}
data class Pair1894<A, B>(val first: A, val second: B)

fun combine1894(left: Int, right: Int): Pair1894<Int, Int> {
    return Pair1894(left + 65, right + 34)
}
object Registry1895 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1895 = "registry-72"
    }
}
class Counter1896 {
    var current: Int by Delegates.observable(79) { _, _, new -> println(new) }

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
object Registry1897 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1897 = "registry-86"
    }
}
class Counter1898 {
    var current: Int by Delegates.observable(93) { _, _, new -> println(new) }

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
