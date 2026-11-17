package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter23700 {
    var current: Int by Delegates.observable(66) { _, _, new -> println(new) }

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
object Registry23701 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23701 = "registry-73"
    }
}
sealed class Outcome23702 {
    data class Success(val payload: String) : Outcome23702()
    data object Failure : Outcome23702()
}

fun describe23702(outcome: Outcome23702): String = when (outcome) {
    is Outcome23702.Success -> outcome.payload
    Outcome23702.Failure -> "failed"
}
class Counter23703 {
    var current: Int by Delegates.observable(87) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull23704(): T? = if (size > 1) this[1] else null

fun render23704(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull23705(): T? = if (size > 1) this[1] else null

fun render23705(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
