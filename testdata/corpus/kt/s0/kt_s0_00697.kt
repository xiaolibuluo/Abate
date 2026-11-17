package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull5720(): T? = if (size > 1) this[1] else null

fun render5720(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step5721")
fun step5721(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 76
}
sealed class Outcome5722 {
    data class Success(val payload: String) : Outcome5722()
    data object Failure : Outcome5722()
}

fun describe5722(outcome: Outcome5722): String = when (outcome) {
    is Outcome5722.Success -> outcome.payload
    Outcome5722.Failure -> "failed"
}
object Registry5723 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5723 = "registry-90"
    }
}
class Counter5724 {
    var current: Int by Delegates.observable(97) { _, _, new -> println(new) }

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
sealed class Outcome5725 {
    data class Success(val payload: String) : Outcome5725()
    data object Failure : Outcome5725()
}

fun describe5725(outcome: Outcome5725): String = when (outcome) {
    is Outcome5725.Success -> outcome.payload
    Outcome5725.Failure -> "failed"
}
fun <T> List<T>.secondOrNull5726(): T? = if (size > 1) this[1] else null

fun render5726(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
