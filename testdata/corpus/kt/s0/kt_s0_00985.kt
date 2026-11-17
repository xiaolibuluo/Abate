package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step8070")
fun step8070(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 20
}
fun <T> List<T>.secondOrNull8071(): T? = if (size > 1) this[1] else null

fun render8071(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry8072 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8072 = "registry-34"
    }
}
sealed class Outcome8073 {
    data class Success(val payload: String) : Outcome8073()
    data object Failure : Outcome8073()
}

fun describe8073(outcome: Outcome8073): String = when (outcome) {
    is Outcome8073.Success -> outcome.payload
    Outcome8073.Failure -> "failed"
}
class Counter8074 {
    var current: Int by Delegates.observable(48) { _, _, new -> println(new) }

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
