package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter19430 {
    var current: Int by Delegates.observable(44) { _, _, new -> println(new) }

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
object Registry19431 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19431 = "registry-51"
    }
}
sealed class Outcome19432 {
    data class Success(val payload: String) : Outcome19432()
    data object Failure : Outcome19432()
}

fun describe19432(outcome: Outcome19432): String = when (outcome) {
    is Outcome19432.Success -> outcome.payload
    Outcome19432.Failure -> "failed"
}
class Counter19433 {
    var current: Int by Delegates.observable(65) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull19434(): T? = if (size > 1) this[1] else null

fun render19434(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull19435(): T? = if (size > 1) this[1] else null

fun render19435(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome19436 {
    data class Success(val payload: String) : Outcome19436()
    data object Failure : Outcome19436()
}

fun describe19436(outcome: Outcome19436): String = when (outcome) {
    is Outcome19436.Success -> outcome.payload
    Outcome19436.Failure -> "failed"
}
