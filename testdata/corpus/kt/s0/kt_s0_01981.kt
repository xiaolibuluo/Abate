package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry16370 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16370 = "registry-43"
    }
}
sealed class Outcome16371 {
    data class Success(val payload: String) : Outcome16371()
    data object Failure : Outcome16371()
}

fun describe16371(outcome: Outcome16371): String = when (outcome) {
    is Outcome16371.Success -> outcome.payload
    Outcome16371.Failure -> "failed"
}
@JvmName("step16372")
fun step16372(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
class Counter16373 {
    var current: Int by Delegates.observable(64) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull16374(): T? = if (size > 1) this[1] else null

fun render16374(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull16375(): T? = if (size > 1) this[1] else null

fun render16375(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
