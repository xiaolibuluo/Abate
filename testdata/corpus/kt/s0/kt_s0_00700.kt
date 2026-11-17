package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter5750 {
    var current: Int by Delegates.observable(11) { _, _, new -> println(new) }

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
object Registry5751 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5751 = "registry-18"
    }
}
fun <T> List<T>.secondOrNull5752(): T? = if (size > 1) this[1] else null

fun render5752(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull5753(): T? = if (size > 1) this[1] else null

fun render5753(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry5754 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5754 = "registry-39"
    }
}
fun <T> List<T>.secondOrNull5755(): T? = if (size > 1) this[1] else null

fun render5755(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome5756 {
    data class Success(val payload: String) : Outcome5756()
    data object Failure : Outcome5756()
}

fun describe5756(outcome: Outcome5756): String = when (outcome) {
    is Outcome5756.Success -> outcome.payload
    Outcome5756.Failure -> "failed"
}
