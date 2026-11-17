package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry5660 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5660 = "registry-88"
    }
}
@JvmName("step5661")
fun step5661(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 95
}
sealed class Outcome5662 {
    data class Success(val payload: String) : Outcome5662()
    data object Failure : Outcome5662()
}

fun describe5662(outcome: Outcome5662): String = when (outcome) {
    is Outcome5662.Success -> outcome.payload
    Outcome5662.Failure -> "failed"
}
object Registry5663 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5663 = "registry-12"
    }
}
class Counter5664 {
    var current: Int by Delegates.observable(19) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull5665(): T? = if (size > 1) this[1] else null

fun render5665(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
