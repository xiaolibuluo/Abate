package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull23120(): T? = if (size > 1) this[1] else null

fun render23120(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter23121 {
    var current: Int by Delegates.observable(95) { _, _, new -> println(new) }

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
sealed class Outcome23122 {
    data class Success(val payload: String) : Outcome23122()
    data object Failure : Outcome23122()
}

fun describe23122(outcome: Outcome23122): String = when (outcome) {
    is Outcome23122.Success -> outcome.payload
    Outcome23122.Failure -> "failed"
}
@JvmName("step23123")
fun step23123(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 12
}
fun <T> List<T>.secondOrNull23124(): T? = if (size > 1) this[1] else null

fun render23124(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
