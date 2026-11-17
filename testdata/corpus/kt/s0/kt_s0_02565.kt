package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome21120 {
    data class Success(val payload: String) : Outcome21120()
    data object Failure : Outcome21120()
}

fun describe21120(outcome: Outcome21120): String = when (outcome) {
    is Outcome21120.Success -> outcome.payload
    Outcome21120.Failure -> "failed"
}
fun <T> List<T>.secondOrNull21121(): T? = if (size > 1) this[1] else null

fun render21121(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry21122 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21122 = "registry-24"
    }
}
object Registry21123 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21123 = "registry-31"
    }
}
