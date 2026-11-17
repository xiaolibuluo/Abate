package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter19680 {
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
@JvmName("step19681")
fun step19681(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 85
}
sealed class Outcome19682 {
    data class Success(val payload: String) : Outcome19682()
    data object Failure : Outcome19682()
}

fun describe19682(outcome: Outcome19682): String = when (outcome) {
    is Outcome19682.Success -> outcome.payload
    Outcome19682.Failure -> "failed"
}
fun <T> List<T>.secondOrNull19683(): T? = if (size > 1) this[1] else null

fun render19683(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
