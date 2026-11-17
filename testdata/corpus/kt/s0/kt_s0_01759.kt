package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome14420 {
    data class Success(val payload: String) : Outcome14420()
    data object Failure : Outcome14420()
}

fun describe14420(outcome: Outcome14420): String = when (outcome) {
    is Outcome14420.Success -> outcome.payload
    Outcome14420.Failure -> "failed"
}
object Registry14421 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14421 = "registry-37"
    }
}
sealed class Outcome14422 {
    data class Success(val payload: String) : Outcome14422()
    data object Failure : Outcome14422()
}

fun describe14422(outcome: Outcome14422): String = when (outcome) {
    is Outcome14422.Success -> outcome.payload
    Outcome14422.Failure -> "failed"
}
fun <T> List<T>.secondOrNull14423(): T? = if (size > 1) this[1] else null

fun render14423(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry14424 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14424 = "registry-58"
    }
}
class Counter14425 {
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
