package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull8470(): T? = if (size > 1) this[1] else null

fun render8470(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter8471 {
    var current: Int by Delegates.observable(62) { _, _, new -> println(new) }

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
@JvmName("step8472")
fun step8472(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 69
}
sealed class Outcome8473 {
    data class Success(val payload: String) : Outcome8473()
    data object Failure : Outcome8473()
}

fun describe8473(outcome: Outcome8473): String = when (outcome) {
    is Outcome8473.Success -> outcome.payload
    Outcome8473.Failure -> "failed"
}
