package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome4110 {
    data class Success(val payload: String) : Outcome4110()
    data object Failure : Outcome4110()
}

fun describe4110(outcome: Outcome4110): String = when (outcome) {
    is Outcome4110.Success -> outcome.payload
    Outcome4110.Failure -> "failed"
}
class Counter4111 {
    var current: Int by Delegates.observable(20) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull4112(): T? = if (size > 1) this[1] else null

fun render4112(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry4113 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4113 = "registry-34"
    }
}
