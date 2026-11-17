package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull9420(): T? = if (size > 1) this[1] else null

fun render9420(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome9421 {
    data class Success(val payload: String) : Outcome9421()
    data object Failure : Outcome9421()
}

fun describe9421(outcome: Outcome9421): String = when (outcome) {
    is Outcome9421.Success -> outcome.payload
    Outcome9421.Failure -> "failed"
}
object Registry9422 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9422 = "registry-43"
    }
}
sealed class Outcome9423 {
    data class Success(val payload: String) : Outcome9423()
    data object Failure : Outcome9423()
}

fun describe9423(outcome: Outcome9423): String = when (outcome) {
    is Outcome9423.Success -> outcome.payload
    Outcome9423.Failure -> "failed"
}
