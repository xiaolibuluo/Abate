package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull4570(): T? = if (size > 1) this[1] else null

fun render4570(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry4571 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4571 = "registry-36"
    }
}
class Counter4572 {
    var current: Int by Delegates.observable(43) { _, _, new -> println(new) }

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
sealed class Outcome4573 {
    data class Success(val payload: String) : Outcome4573()
    data object Failure : Outcome4573()
}

fun describe4573(outcome: Outcome4573): String = when (outcome) {
    is Outcome4573.Success -> outcome.payload
    Outcome4573.Failure -> "failed"
}
