package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull5240(): T? = if (size > 1) this[1] else null

fun render5240(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter5241 {
    var current: Int by Delegates.observable(34) { _, _, new -> println(new) }

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
sealed class Outcome5242 {
    data class Success(val payload: String) : Outcome5242()
    data object Failure : Outcome5242()
}

fun describe5242(outcome: Outcome5242): String = when (outcome) {
    is Outcome5242.Success -> outcome.payload
    Outcome5242.Failure -> "failed"
}
sealed class Outcome5243 {
    data class Success(val payload: String) : Outcome5243()
    data object Failure : Outcome5243()
}

fun describe5243(outcome: Outcome5243): String = when (outcome) {
    is Outcome5243.Success -> outcome.payload
    Outcome5243.Failure -> "failed"
}
object Registry5244 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5244 = "registry-55"
    }
}
sealed class Outcome5245 {
    data class Success(val payload: String) : Outcome5245()
    data object Failure : Outcome5245()
}

fun describe5245(outcome: Outcome5245): String = when (outcome) {
    is Outcome5245.Success -> outcome.payload
    Outcome5245.Failure -> "failed"
}
