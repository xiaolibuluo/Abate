package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24100 {
    data class Success(val payload: String) : Outcome24100()
    data object Failure : Outcome24100()
}

fun describe24100(outcome: Outcome24100): String = when (outcome) {
    is Outcome24100.Success -> outcome.payload
    Outcome24100.Failure -> "failed"
}
fun <T> List<T>.secondOrNull24101(): T? = if (size > 1) this[1] else null

fun render24101(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry24102 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24102 = "registry-18"
    }
}
fun <T> List<T>.secondOrNull24103(): T? = if (size > 1) this[1] else null

fun render24103(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter24104 {
    var current: Int by Delegates.observable(32) { _, _, new -> println(new) }

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
