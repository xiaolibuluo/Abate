package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry16050 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16050 = "registry-15"
    }
}
class Counter16051 {
    var current: Int by Delegates.observable(22) { _, _, new -> println(new) }

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
object Registry16052 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16052 = "registry-29"
    }
}
object Registry16053 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16053 = "registry-36"
    }
}
@JvmName("step16054")
fun step16054(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 43
}
object Registry16055 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16055 = "registry-50"
    }
}
fun <T> List<T>.secondOrNull16056(): T? = if (size > 1) this[1] else null

fun render16056(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome16057 {
    data class Success(val payload: String) : Outcome16057()
    data object Failure : Outcome16057()
}

fun describe16057(outcome: Outcome16057): String = when (outcome) {
    is Outcome16057.Success -> outcome.payload
    Outcome16057.Failure -> "failed"
}
sealed class Outcome16058 {
    data class Success(val payload: String) : Outcome16058()
    data object Failure : Outcome16058()
}

fun describe16058(outcome: Outcome16058): String = when (outcome) {
    is Outcome16058.Success -> outcome.payload
    Outcome16058.Failure -> "failed"
}
