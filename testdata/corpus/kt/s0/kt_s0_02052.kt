package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome16960 {
    data class Success(val payload: String) : Outcome16960()
    data object Failure : Outcome16960()
}

fun describe16960(outcome: Outcome16960): String = when (outcome) {
    is Outcome16960.Success -> outcome.payload
    Outcome16960.Failure -> "failed"
}
@JvmName("step16961")
fun step16961(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 41
}
fun <T, R> List<T>.mapChecked16962(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter16963 {
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
data class Pair16964<A, B>(val first: A, val second: B)

fun combine16964(left: Int, right: Int): Pair16964<Int, Int> {
    return Pair16964(left + 62, right + 30)
}
fun <T, R> List<T>.mapChecked16965(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair16966<A, B>(val first: A, val second: B)

fun combine16966(left: Int, right: Int): Pair16966<Int, Int> {
    return Pair16966(left + 76, right + 52)
}
object Registry16967 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16967 = "registry-83"
    }
}
