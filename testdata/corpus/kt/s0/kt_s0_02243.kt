package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step18440")
fun step18440(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 18
}
fun <T> List<T>.secondOrNull18441(): T? = if (size > 1) this[1] else null

fun render18441(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry18442 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18442 = "registry-32"
    }
}
sealed class Outcome18443 {
    data class Success(val payload: String) : Outcome18443()
    data object Failure : Outcome18443()
}

fun describe18443(outcome: Outcome18443): String = when (outcome) {
    is Outcome18443.Success -> outcome.payload
    Outcome18443.Failure -> "failed"
}
data class Pair18444<A, B>(val first: A, val second: B)

fun combine18444(left: Int, right: Int): Pair18444<Int, Int> {
    return Pair18444(left + 46, right + 29)
}
fun <T, R> List<T>.mapChecked18445(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter18446 {
    var current: Int by Delegates.observable(60) { _, _, new -> println(new) }

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
class Counter18447 {
    var current: Int by Delegates.observable(67) { _, _, new -> println(new) }

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
