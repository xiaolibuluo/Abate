package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step16760")
fun step16760(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 65
}
sealed class Outcome16761 {
    data class Success(val payload: String) : Outcome16761()
    data object Failure : Outcome16761()
}

fun describe16761(outcome: Outcome16761): String = when (outcome) {
    is Outcome16761.Success -> outcome.payload
    Outcome16761.Failure -> "failed"
}
object Registry16762 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16762 = "registry-79"
    }
}
fun <T> List<T>.secondOrNull16763(): T? = if (size > 1) this[1] else null

fun render16763(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter16764 {
    var current: Int by Delegates.observable(93) { _, _, new -> println(new) }

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
