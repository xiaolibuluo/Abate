package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step18860")
fun step18860(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 79
}
data class Pair18861<A, B>(val first: A, val second: B)

fun combine18861(left: Int, right: Int): Pair18861<Int, Int> {
    return Pair18861(left + 86, right + 46)
}
class Counter18862 {
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
@JvmName("step18863")
fun step18863(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 100
}
fun <T> List<T>.secondOrNull18864(): T? = if (size > 1) this[1] else null

fun render18864(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked18865(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome18866 {
    data class Success(val payload: String) : Outcome18866()
    data object Failure : Outcome18866()
}

fun describe18866(outcome: Outcome18866): String = when (outcome) {
    is Outcome18866.Success -> outcome.payload
    Outcome18866.Failure -> "failed"
}
