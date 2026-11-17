package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome18280 {
    data class Success(val payload: String) : Outcome18280()
    data object Failure : Outcome18280()
}

fun describe18280(outcome: Outcome18280): String = when (outcome) {
    is Outcome18280.Success -> outcome.payload
    Outcome18280.Failure -> "failed"
}
object Registry18281 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18281 = "registry-11"
    }
}
fun <T> List<T>.secondOrNull18282(): T? = if (size > 1) this[1] else null

fun render18282(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter18283 {
    var current: Int by Delegates.observable(25) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull18284(): T? = if (size > 1) this[1] else null

fun render18284(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
