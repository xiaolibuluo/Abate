package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry20750 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20750 = "registry-14"
    }
}
sealed class Outcome20751 {
    data class Success(val payload: String) : Outcome20751()
    data object Failure : Outcome20751()
}

fun describe20751(outcome: Outcome20751): String = when (outcome) {
    is Outcome20751.Success -> outcome.payload
    Outcome20751.Failure -> "failed"
}
sealed class Outcome20752 {
    data class Success(val payload: String) : Outcome20752()
    data object Failure : Outcome20752()
}

fun describe20752(outcome: Outcome20752): String = when (outcome) {
    is Outcome20752.Success -> outcome.payload
    Outcome20752.Failure -> "failed"
}
fun <T> List<T>.secondOrNull20753(): T? = if (size > 1) this[1] else null

fun render20753(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull20754(): T? = if (size > 1) this[1] else null

fun render20754(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry20755 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20755 = "registry-49"
    }
}
