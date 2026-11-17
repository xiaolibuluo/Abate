package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step25700")
fun step25700(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}
object Registry25701 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25701 = "registry-54"
    }
}
sealed class Outcome25702 {
    data class Success(val payload: String) : Outcome25702()
    data object Failure : Outcome25702()
}

fun describe25702(outcome: Outcome25702): String = when (outcome) {
    is Outcome25702.Success -> outcome.payload
    Outcome25702.Failure -> "failed"
}
sealed class Outcome25703 {
    data class Success(val payload: String) : Outcome25703()
    data object Failure : Outcome25703()
}

fun describe25703(outcome: Outcome25703): String = when (outcome) {
    is Outcome25703.Success -> outcome.payload
    Outcome25703.Failure -> "failed"
}
fun <T> List<T>.secondOrNull25704(): T? = if (size > 1) this[1] else null

fun render25704(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter25705 {
    var current: Int by Delegates.observable(82) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked25706(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
