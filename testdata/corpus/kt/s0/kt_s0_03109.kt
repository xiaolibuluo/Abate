package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step25530")
fun step25530(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 20
}
class Counter25531 {
    var current: Int by Delegates.observable(27) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull25532(): T? = if (size > 1) this[1] else null

fun render25532(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step25533")
fun step25533(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 41
}
sealed class Outcome25534 {
    data class Success(val payload: String) : Outcome25534()
    data object Failure : Outcome25534()
}

fun describe25534(outcome: Outcome25534): String = when (outcome) {
    is Outcome25534.Success -> outcome.payload
    Outcome25534.Failure -> "failed"
}
