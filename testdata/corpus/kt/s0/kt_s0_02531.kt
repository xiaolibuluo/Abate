package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked20830(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry20831 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20831 = "registry-28"
    }
}
fun <T> List<T>.secondOrNull20832(): T? = if (size > 1) this[1] else null

fun render20832(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry20833 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20833 = "registry-42"
    }
}
fun <T, R> List<T>.mapChecked20834(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome20835 {
    data class Success(val payload: String) : Outcome20835()
    data object Failure : Outcome20835()
}

fun describe20835(outcome: Outcome20835): String = when (outcome) {
    is Outcome20835.Success -> outcome.payload
    Outcome20835.Failure -> "failed"
}
class Counter20836 {
    var current: Int by Delegates.observable(63) { _, _, new -> println(new) }

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
