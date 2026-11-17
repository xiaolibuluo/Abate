package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter13000 {
    var current: Int by Delegates.observable(27) { _, _, new -> println(new) }

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
object Registry13001 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13001 = "registry-34"
    }
}
@JvmName("step13002")
fun step13002(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 41
}
sealed class Outcome13003 {
    data class Success(val payload: String) : Outcome13003()
    data object Failure : Outcome13003()
}

fun describe13003(outcome: Outcome13003): String = when (outcome) {
    is Outcome13003.Success -> outcome.payload
    Outcome13003.Failure -> "failed"
}
fun <T> List<T>.secondOrNull13004(): T? = if (size > 1) this[1] else null

fun render13004(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry13005 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13005 = "registry-62"
    }
}
sealed class Outcome13006 {
    data class Success(val payload: String) : Outcome13006()
    data object Failure : Outcome13006()
}

fun describe13006(outcome: Outcome13006): String = when (outcome) {
    is Outcome13006.Success -> outcome.payload
    Outcome13006.Failure -> "failed"
}
