package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull22200(): T? = if (size > 1) this[1] else null

fun render22200(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter22201 {
    var current: Int by Delegates.observable(63) { _, _, new -> println(new) }

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
object Registry22202 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22202 = "registry-70"
    }
}
sealed class Outcome22203 {
    data class Success(val payload: String) : Outcome22203()
    data object Failure : Outcome22203()
}

fun describe22203(outcome: Outcome22203): String = when (outcome) {
    is Outcome22203.Success -> outcome.payload
    Outcome22203.Failure -> "failed"
}
fun <T> List<T>.secondOrNull22204(): T? = if (size > 1) this[1] else null

fun render22204(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
