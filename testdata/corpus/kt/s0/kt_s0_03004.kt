package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry24700 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24700 = "registry-8"
    }
}
object Registry24701 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24701 = "registry-15"
    }
}
sealed class Outcome24702 {
    data class Success(val payload: String) : Outcome24702()
    data object Failure : Outcome24702()
}

fun describe24702(outcome: Outcome24702): String = when (outcome) {
    is Outcome24702.Success -> outcome.payload
    Outcome24702.Failure -> "failed"
}
sealed class Outcome24703 {
    data class Success(val payload: String) : Outcome24703()
    data object Failure : Outcome24703()
}

fun describe24703(outcome: Outcome24703): String = when (outcome) {
    is Outcome24703.Success -> outcome.payload
    Outcome24703.Failure -> "failed"
}
class Counter24704 {
    var current: Int by Delegates.observable(36) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull24705(): T? = if (size > 1) this[1] else null

fun render24705(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
