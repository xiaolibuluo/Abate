package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair21750<A, B>(val first: A, val second: B)

fun combine21750(left: Int, right: Int): Pair21750<Int, Int> {
    return Pair21750(left + 53, right + 45)
}
fun <T> List<T>.secondOrNull21751(): T? = if (size > 1) this[1] else null

fun render21751(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry21752 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21752 = "registry-67"
    }
}
fun <T, R> List<T>.mapChecked21753(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull21754(): T? = if (size > 1) this[1] else null

fun render21754(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry21755 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21755 = "registry-88"
    }
}
sealed class Outcome21756 {
    data class Success(val payload: String) : Outcome21756()
    data object Failure : Outcome21756()
}

fun describe21756(outcome: Outcome21756): String = when (outcome) {
    is Outcome21756.Success -> outcome.payload
    Outcome21756.Failure -> "failed"
}
