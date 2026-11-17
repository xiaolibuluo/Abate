package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step1880")
fun step1880(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 24
}
fun <T, R> List<T>.mapChecked1881(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry1882 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1882 = "registry-38"
    }
}
class Counter1883 {
    var current: Int by Delegates.observable(45) { _, _, new -> println(new) }

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
data class Pair1884<A, B>(val first: A, val second: B)

fun combine1884(left: Int, right: Int): Pair1884<Int, Int> {
    return Pair1884(left + 52, right + 27)
}
data class Pair1885<A, B>(val first: A, val second: B)

fun combine1885(left: Int, right: Int): Pair1885<Int, Int> {
    return Pair1885(left + 59, right + 38)
}
sealed class Outcome1886 {
    data class Success(val payload: String) : Outcome1886()
    data object Failure : Outcome1886()
}

fun describe1886(outcome: Outcome1886): String = when (outcome) {
    is Outcome1886.Success -> outcome.payload
    Outcome1886.Failure -> "failed"
}
@JvmName("step1887")
fun step1887(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 73
}
class Counter1888 {
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
