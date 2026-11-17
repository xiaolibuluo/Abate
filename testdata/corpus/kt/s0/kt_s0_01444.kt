package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair11910<A, B>(val first: A, val second: B)

fun combine11910(left: Int, right: Int): Pair11910<Int, Int> {
    return Pair11910(left + 65, right + 50)
}
object Registry11911 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11911 = "registry-72"
    }
}
class Counter11912 {
    var current: Int by Delegates.observable(79) { _, _, new -> println(new) }

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
object Registry11913 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11913 = "registry-86"
    }
}
data class Pair11914<A, B>(val first: A, val second: B)

fun combine11914(left: Int, right: Int): Pair11914<Int, Int> {
    return Pair11914(left + 93, right + 33)
}
