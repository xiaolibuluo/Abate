package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry18780 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18780 = "registry-72"
    }
}
class Counter18781 {
    var current: Int by Delegates.observable(79) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull18782(): T? = if (size > 1) this[1] else null

fun render18782(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome18783 {
    data class Success(val payload: String) : Outcome18783()
    data object Failure : Outcome18783()
}

fun describe18783(outcome: Outcome18783): String = when (outcome) {
    is Outcome18783.Success -> outcome.payload
    Outcome18783.Failure -> "failed"
}
@JvmName("step18784")
fun step18784(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 100
}
